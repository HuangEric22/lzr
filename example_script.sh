#sudo zmap --seed=88 --target-port=443 -r 750000 --output-filter="success = 1 && repeat = 0"  -f"saddr,daddr,sport,dport,seqnum,acknum,window" -O json --source-ip="171.67.71.98" -i ens8 |  sudo ./lzr -t 5 -w 6 -h tls,http -f /mnt/projects/continuous_scanning/LZR/f_LZR_ZGrab443_TLSHTTP.json | ztee /mnt/projects/continuous_scanning/LZR/f_LZR_ztee_ZGrab443_TLSHTTP.json |  sudo /home/lizhikev/go/src/github.com/zmap/zgrab2/./zgrab2 multiple -c multiple.ini -o /mnt/projects/continuous_scanning/LZR/LZR_zgrab443_TLSHTTP.json -s 5000
#sudo zmap --seed=88  --target-port=443 -r 750000 --output-filter="success = 1 && repeat = 0" | ztee /mnt/projects/continuous_scanning/LZR/zgrab443_ztee_TLSHTTP.json | sudo /home/lizhikev/go/src/github.com/zmap/zgrab2/./zgrab2 multiple -c multiple.ini -o /mnt/projects/continuous_scanning/LZR/zgrab443_TLSHTTP.json -s 5000

#sudo zmap --seed=88 --target-port=443 -r 750000 --output-filter="success = 1 && repeat = 0"  -f"saddr,daddr,sport,dport,seqnum,acknum,window" -O json --source-ip="171.67.71.98" -i ens8 |  sudo ./lzr -t 5 -w 6 -h tls -f /mnt/projects/continuous_scanning/LZR/f_LZR_cpuRun.json -cpuprofile lzr.prof

sudo  zmap  --target-port=443 -r 2000000 --output-filter="success = 1 && repeat = 0"  -f"saddr,daddr,sport,dport,seqnum,acknum,window" -O json --source-ip="171.67.71.98" -i ens8 |  sudo ./lzr -t 3 -w 6 -h tls,http -cpuprofile lzr_full.prof -f /mnt/projects/continuous_scanning/LZR/fingerprint443_TLSHTTP_CPU.json4 > /mnt/projects/continuous_scanning/LZR/fingerprint_out_4