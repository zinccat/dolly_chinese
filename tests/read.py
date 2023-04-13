import jsonlines

with open("test.jsonl", "r+", encoding="utf-8") as f:
    for item in jsonlines.Reader(f):
        print(item)