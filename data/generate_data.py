import json
import argparse
import subprocess

# parameters
parser = argparse.ArgumentParser()

parser.add_argument(dest="event_name")
args = parser.parse_args()

event_name = args.event_name

while True:
    result = subprocess.getstatusoutput([f"synth generate ./events/{event_name}"])[1]
    result = json.loads(result)