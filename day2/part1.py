

def main():
    with open('input.txt','r') as f:
        p=f.readlines()
    p=[line[:-1] for line in p]
    #print(p)
    total = 0 
    for key,value in enumerate(p):
        sides = []

        for dims in value.split("x"):
            sides.append(dims)
        #print(sides)
        
        sides.sort()
        
        total += int(sides[0])*int(sides[1])*3
        total += int(sides[0])*int(sides[2])*2
        total += int(sides[1])*int(sides[2])*2

    print(total)




main()
