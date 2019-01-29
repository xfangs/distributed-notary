#!/usr/bin/env bash
# 删除历史数据
#rm -rf .notary*
#rm -f .n*.txt
# 启动两条测试链
#cd ../deploygeth
#sh ./deploygeth.sh
#cd ../run
#kill 历史dnotary
ps -ef | grep dnotary  | grep -v grep | awk '{print $2}' |xargs kill -9
#0
dnotary --address=0x1a9ec3b0b807464e6d3398a59d6b0a369bf422fa --user-listen=127.0.0.1:3330 --notary-listen=127.0.0.1:33300 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n0  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n0.txt 2>&1 &
#1
dnotary --address=0x33df901abc22dcb7f33c2a77ad43cc98fbfa0790 --user-listen=127.0.0.1:3331 --notary-listen=127.0.0.1:33301 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n1  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n1.txt 2>&1 &
#2
dnotary --address=0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401 --user-listen=127.0.0.1:3332 --notary-listen=127.0.0.1:33302 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n2  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n2.txt 2>&1 &
#3
dnotary --address=0xc4c08f9227be0f1750f5d5467eed462ec133b15e --user-listen=127.0.0.1:3333 --notary-listen=127.0.0.1:33303 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n3  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n3.txt 2>&1 &
#4
dnotary --address=0x543fc024cdd1f0d346a306f5e99ec0d8fe392920 --user-listen=127.0.0.1:3334 --notary-listen=127.0.0.1:33304 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n4  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n4.txt 2>&1 &
#5
dnotary --address=0x920a90acc9164272ede4ae1e9c33841f019f53a4 --user-listen=127.0.0.1:3335 --notary-listen=127.0.0.1:33305 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n5  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n5.txt 2>&1 &
#6
dnotary --address=0x215c0d259ac31571a43295f2e411a697cd30748c --user-listen=127.0.0.1:3336 --notary-listen=127.0.0.1:33306 --notary-config-file=./notary.conf --keystore-path=../keystore --datadir=./.notary_n6  --smc-rpc-point="http://127.0.0.1:17888" --eth-rpc-point="http://127.0.0.1:19888"  >>.n6.txt 2>&1 &



