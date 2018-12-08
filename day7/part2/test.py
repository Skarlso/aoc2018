#!/usr/bin/env python3

import re

INSTRUCTION_PATTERN = r'^Step (\w) must be finished before step (\w) can begin.$'
NUMBER_OF_WORKERS = 5
TASK_DEFAULT_LENGTH = 60

with open('input.txt') as f:
    raw_instructions = [line for line in f.read().split('\n') if line != '']

task_graph = {}
list_of_tasks = []

for line in raw_instructions:
    (task1, task2) = re.search(INSTRUCTION_PATTERN, line).groups()
    task_graph[task2] = task_graph.get(task2, []) + [task1]
    list_of_tasks.append(task1)
    list_of_tasks.append(task2)

list_of_tasks = sorted(set(list_of_tasks))
final_sequence = []

for add in list_of_tasks:
    if add not in task_graph:
        task_graph[add] = []

def done(task):
    final_sequence.append(task)
    for t in task_graph:
        if task_graph[t] is not None and task in task_graph[t]:
            task_graph[t].remove(task)

class Worker:
    item = None
    busy_until = -1

    def is_busy(self, time):
        return time <= self.busy_until

    def is_ready(self, time):
        return self.item is not None and time >= self.busy_until

    def done(self):
        print(f'Task "{self.item}" is done [{id(self)}]')
        task = self.item
        self.item = None
        return task

    def assign(self, task, current_time):
        task_graph[task] = None
        self.item = task
        self.busy_until = current_time + TASK_DEFAULT_LENGTH + (ord(task) - 64) - 1
        print(f'Task "{task}" assigned to {id(self)} -> {self.busy_until}')

workers = [Worker() for _ in range(0, NUMBER_OF_WORKERS)]

current_time = 0
while len(final_sequence) < len(list_of_tasks):
    #print(f'>>> {current_time}s   -> {final_sequence}')

    can_we_start = []
    for task in list_of_tasks:
        deps = task_graph.get(task, None)
        if deps is not None and len(deps) < 1:
            can_we_start.append(task)

    [done(worker.done()) for worker in workers if worker.is_ready(current_time)]

    #can_we_start.sort()
    #print(current_time, can_we_start)

    for task in can_we_start:
        for worker in workers:
            if not worker.is_busy(current_time):
                worker.assign(task, current_time)
                break

    current_time += 1

print(f"{current_time}s")
print("".join(final_sequence))