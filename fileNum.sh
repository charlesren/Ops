#/bin/sh
export LC_ALL="en_US"
export LANG="en_US"
# ##################################################################################################
# Description: to check informix tablespace state check
# Parameters: INI file nane, HOSTIP (number) and HOSTIP (dot)
# Analysis Rule: STATUS|CHECKTIME|VALUE|THRESHOLD|DIRNAME|DIRSIZE|FILETYPE
# Author: Rencc
# Date: 2019.09.11
# VERSION:1
# ##################################################################################################
RPath=`pwd`
#	load common source
. ${RPath}/scripts/base/comFunctions.sh
# #####################################
#	Main Program
# #####################################
#脚本引用的ini文件名称
INIFILE=${1}

#接收主机IP
HP12=${2}

#主机IP(xx.yy.z.nn)
HOSTIP=${3}

#获得脚本名称
scriptsName=`echo ${0##*/} | cut -f 1 -d "."`
programName="${0##*/}"
#检查时间
vChkTime=`date "+%F %T"`

TMP="${RPath}/temp"
OUT="${RPath}/out"
LOG="${RPath}/log"

TMPFILE="${TMP}/${scriptsName}${HP12}.out"
LOGFILE="${LOG}/${scriptsName}${HP12}.log"
OUTFILE="${OUT}/check${HP12}.out"

> ${TMPFILE}
> ${LOGFILE}
while read vR

do

done < ${INIFILE}
cat ${TMPFILE} >> ${OUTFILE}