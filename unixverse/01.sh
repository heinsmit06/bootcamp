#!/bin/bash
awk 'BEGIN{FS=OFS=","} {tmp2 = $2; tmp4 = $4 ; $2=tmp4; $4=tmp2}1' movies.csv
