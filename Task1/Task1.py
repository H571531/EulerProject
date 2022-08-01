import timeit
import argparse

parser = argparse.ArgumentParser(description='Sum of 3 and 5 multiplications')
parser.add_argument('-n',required=True)
args = parser.parse_args()

start = timeit.default_timer()
sum=0
for i in range(int(args.n)):
    if(i%3==0 or i%5==0):
        sum=sum+i
stop = timeit.default_timer()
print('Time: ', stop - start)  
print(sum)