from tqdm import tqdm
from utils import translate
import jsonlines

idx = 0 # selecting the idx's part of dataset
split = list(range(0, 15000, 100)) + [14999]

cnt = 0

category = []

with open("databricks-dolly-15k.jsonl", "r+", encoding="utf-8") as f:
    data = []
    for item in jsonlines.Reader(f):
        if cnt >= split[idx] and cnt < split[idx+1]:
            data.append(item['instruction'])
            data.append(item['context'])
            data.append(item['response'])
            category.append(item['category'])
        elif cnt >= split[idx+1]:
            break
        cnt += 1

# need parallel
ret = [translate(d) for d in tqdm(data)]

instruction_tr = ret[0::3]
context_tr = ret[1::3]
response_tr = ret[2::3]

with open("dolly_chinese_{}.jsonl".format(idx), "w", encoding="utf-8") as f:
    writer = jsonlines.Writer(f)
    for i in range(len(instruction_tr)):
        writer.write({'instruction': instruction_tr[i], 'context': context_tr[i], 'response': response_tr[i], 'category': category[i]})