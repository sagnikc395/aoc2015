

def main():
    with open('input.txt','r') as f:
        p=f.readlines()

    p=[line[:-1] for line in p]
    #print(p)
    total = 0 
    sides = []

    for key,value in enumerate(p):
        res = value.split("x")
        sides.append(res)

    sides.sort()

    total += (sides[0]+sides[1])*2
    total += sides[0]*sides[1]*sides[2]
    
    print(total)



main()
