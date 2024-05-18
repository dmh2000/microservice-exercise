
# Import module
import os
 
# Assign directory
directory = os.getcwd()

def subdir(directory):
    os.system(f"echo {directory} && cd {directory} && go mod tidy && cd ..")

    for path, folders, files in os.walk(directory):
        # List contain of folder
        # print(path)
        for folder_name in folders:
            subdir(f"{directory}/{folder_name}")


if __name__ == '__main__':
    subdir(directory)
