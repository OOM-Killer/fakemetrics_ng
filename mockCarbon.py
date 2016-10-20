#!/usr/bin/env python


import sys
import signal
import SocketServer
import re
import collections as c


class State(object):

    def __init__(self):
        self.invalid_lines = 0
        self.records = 0
        self.key_counts = c.defaultdict(int)
        self.key_values = c.defaultdict(list)
        self.times = []

    def invalid_format(self):
        self.invalid_lines += 1

    def record(self, key, value, time):
        self.key_counts[key] += 1
        self.key_values[key].append(value)
        self.times.append(time)

    def print_state(self):
        print('key count: {v}'.format(v=len(self.key_counts)))
        print(
            'metrics per key avg: {v}'
            .format(v=sum(self.key_counts.values()) / len(self.key_counts)),
        )
        print(
            'metrics per key max: {v}'
            .format(v=max(self.key_counts.values())),
        )


class MockCarbonHandler(SocketServer.StreamRequestHandler):
    f = re.compile('^([^ ]+) ([0-9\.]+) ([0-9\.]+)$')

    def handle(self):
        line = self.rfile.readline().strip()
        if not self.f.match(line):
            state.invalid_format()
            return

        data = map(
            lambda x: x.strip(),
            self.f.findall(line)[0]
        )

        state.record(data[0], *[float(x) for x in data[1:]])


if __name__ == '__main__':
    state = State()
    server = SocketServer.TCPServer(
        (
            '127.0.0.1',
            2003,
        ),
        MockCarbonHandler,
    )

    def term_handler(*args, **kwargs):
        state.print_state()
        sys.exit(0)

    signal.signal(signal.SIGINT, term_handler)
    server.serve_forever()
