import os, json
import subprocess
from collections import defaultdict

# Read the data from the JSON file.
print("Reading data from JSON file...")
with open('data.json', 'r') as json_file:
    json_list = list(json_file)

datas = []

for json_str in json_list:
    result = json.loads(json_str)
    datas.append(result)

# Organize the data.
print("Organizing data...")
organized_datas = defaultdict(list)

for data in datas:
    event_name = data['event_name']
    organized_datas[event_name].append(data)

# Write the data to a file.
with open('organized_data.json', 'w') as outfile:
    json.dump(organized_datas, outfile, ensure_ascii=False, indent=4)

# Run synth
print("Writing data to a file...")
process = subprocess.Popen(["synth import events/ --from json:organized_data.json"], shell=True)
process.wait()

# Transform the data to a valid JSON format.
print("Transforming data to a valid JSON format...")
for key in organized_datas.keys():
    with open(f'./events/{key}.json', 'r') as outfile:
        data = json.load(outfile)

    data["type"] = "object"
    del data["length"]

    os.mkdir(f'./events/{key}')

    with open(f'./events/{key}/{key}.json', 'w') as outfile:
        json.dump(data, outfile, ensure_ascii=False, indent=4)

    os.remove(f'./events/{key}.json')