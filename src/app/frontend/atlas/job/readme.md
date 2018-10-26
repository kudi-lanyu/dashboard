atlas_top

=========================
[mengliang@15-pxe atlasctl]$ atlasctl top
 Display Resource (GPU) usage.

Available Commands:
node        Display Resource (GPU) usage of nodes
job         Display Resource (GPU) usage of pods

Usage:
  atlasctl top [flags]
  atlasctl top [command]

Available Commands:
  job         Display Resource (GPU) usage of jobs.
  node        Display Resource (GPU) usage of nodes.


=====================
[mengliang@15-pxe atlasctl]$ atlasctl top node
NAME               IPADDRESS    ROLE    GPU(Total)  GPU(Allocated)
00-25-90-c0-f7-88  10.10.15.34  master  0           0
00-25-90-c0-f7-c8  10.10.15.98  <none>  0           0
0c-c4-7a-15-e1-9c  10.10.15.5   <none>  0           0
-----------------------------------------------------------------------------------------
Allocated/Total GPUs In Cluster:
0/0 (0%)

==========================
[mengliang@15-pxe atlasctl]$ atlasctl top job
NAME  STATUS     TRAINER  AGE  NODE  GPU(Requests)  GPU(Allocated)
mj    SUCCEEDED  MPIJOB   21h  N/A   0              0


Total Allocated GPUs of Training Job:
0

Total Requested GPUs of Training Job:
0


================================
atlasctl 中的结构体信息
