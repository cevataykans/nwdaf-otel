# nwdaf-otel

Hi Benjamin! This is the configutation I use for the core. After installing the core, I verify that I can bind to the uesimtun0 interface.

Make sure that once you clone my aether-onramp repo, you do it with --recursive option:
```bash
# ssh
git clone --recursive git@github.com:cevataykans/aether-onramp.git
# or https
git clone --recursive https://github.com/cevataykans/aether-onramp.git
```

After this, head over to the hosts.ini and edit it accordingly.

Some tips that I can give is:
* Make sure your user can run "sudo" on both nuc1 and nuc3 without requiring password.
  * While this is dangerous, it may be necessary for some commands ansible may be using.
* You can just copy the installation scripts located under scripts/infra and use them as you see fit.
  * You can also utilize the make command for quick testing.
  * Do not forget to edit the paths in both install/uninstall scripts.
* Make sure the firewall is turned off, 
  * installation script tries to disable it first, but idk if this change requires a system restart for it to take affect.

In aether-onramp, navigate to the vars/main.yml and **only** change:
* core.amf.ip (should be nuc3 IP)
* Thats it, if I have not missed anything!

Install the core with:
```bash
make install
```

After ueransim is installed, and make (or installation script) exits:
* make sure they are not any errors reported.
* If there is any errors, it may stop the execution for the remaining commands run by Ansible.
* Make sure everything is healthy by looking at pods and their states.
  * e.g. there should be no CrashLoopBackoff

Run GNB with:
```bash
sudo ./nr-gnb -c ../config/custom-gnb.yaml
```
Run UE with:
```bash
sudo ./nr-ue -c ../config/custom-ue.yaml
```

> I received an error with ne-ue that I was not running it with sudo (binary required some special privileges). 
> To mitigate this, I run both gnb and ue in sudo mode.

Uninstall the core with:
```bash
make uninstall
```

**ALWAYS MAKE CLEAR UNINSTALL**. If something is manually added, it may not be removed by ansible, causing installation problems in the future. I do not manipulate any ip table or routes, everything is handled automatically by Ansible with these configs.