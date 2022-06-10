# Lectures
Experiments and sandboxes to support teaching at TUBS

## Prerequisites
Most of the subfolders contain vagrant environments for which you need a machine with [Vagrant](https://www.vagrantup.com/docs/installation) and either [KVM](https://help.ubuntu.com/community/KVM/Installation) (recommended on Unix-based systems) or [Virtualbox](https://www.virtualbox.org/) (for Windows and Unix) as a hypervisor. Please understand that I won't be able to provide support for the installation of these, so follow the websites and use the depths of the internet in case of problems.
In some cases, virtualization might not be enabled by default, so you have to [enable it in the BIOS of your pc](https://www.technewstoday.com/how-to-enable-virtualization-in-bios/).

## How it works
Each Vagrantfile creates one or more virtual machines and connects them to each other when you [cd](https://www.wikihow.com/Change-Directories-in-Command-Prompt) into the file's directory on the terminal/cmd and execute "vagrant up". These machines can always be destroyed (and should be once you don't need them anymore) by issuing "vagrant destroy -f" from the same directory. To see if the machines of the current directory are running, you can type "vagrant status". I strongly recommend to make sure they are destroyed before you experiment with another set of machines.
