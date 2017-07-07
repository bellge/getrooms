# getrooms
定时采集一些直播平台的直播房间总数，目前采集的平台包括：触手、飞云、企鹅电竞、大神互动，采集的数据保存到xlsx文件，方便分析。

# Usage:
getrooms   [-i int]	
     -i为可选参数，指定采集数据的时间间隔，默认30分钟采集一次

example:
>./getrooms  or  ./getrooms -i 30
