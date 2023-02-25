#!/bin/bash

# Prompt the user to enter the input file name and input_file is a user define variable
echo -e "Please select a JSON file"
input_file=$(gum filter --fuzzy --prompt="Select JSON file: " --prompt.foreground="4" --match.foreground="1")
echo "You have selected: $input_file"

# Check if the input file is a JSON file
if [ "${input_file##*.}" != "json" ]; then
  echo "Error: input file is not a JSON file."
  exit 1
fi

sleep 1 
# Prompt to choose file format

echo "Choose the output file format: "

choice=$(gum choose "1) .txt" \
                    "2) .csv")

# Define the output file name based on the user's choice
if [ "$choice" == "1) .txt" ]; then
output_file="letters.txt"
delimiter="\t"
format="tab-delimited"
elif [ "$choice" == "2) .csv" ]; then
output_file="letters.csv"
delimiter=","
format="comma-separated"
else
echo "Invalid choice. Exiting..."
exit 1
fi

# Use jq to extract the header row and use the selected delimiter
header_row="No${delimiter}Date${delimiter}In/Out${delimiter}Registration${delimiter}Sender${delimiter}Recipient${delimiter}Filename"
# Use jq to extract the data rows and use the selected delimiter
data_rows=$(jq -r '.[] | [.no, .date, .in_out, .registration, .sender, .recipient, .file_name] | @tsv' "$input_file")

# Check if the data rows are empty
if [ -z "$data_rows" ]; then
echo "Error: data rows are empty"
exit 1
fi

# Replace the tabs in the data rows with the selected delimiter
data_rows=$(echo "$data_rows" | tr '\t' "$delimiter")


#Combine the header row and data rows into the output file
echo -e "$header_row" > "$output_file"
echo -e "$data_rows" >> "$output_file"

gum spin --spinner dot --title "Please wait..." -- sleep 2

#Notify the user that the conversion is complete
echo "The JSON data has been converted to a $format file ($output_file)"

