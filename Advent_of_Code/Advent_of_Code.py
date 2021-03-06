#''' Day 1-1 '''
#input = "892195969991735837915273868729548694237967495115412399373194562526947585337233793568278265279199883197167634791293177986152566236718332617536487236879747167999983363832257912445756887314879229925864477761357139855548522513798899853896612387146687716264599943289416326727256525173953861534244979466587895429399159924916364476319573895566795393368411672387263615582128377676293612892723762237191146714286233543514411813323197995953854871628225358543514157867372265718724276911699514971458844849349726276329135118243155698271218844347387457343656446381799296893222256198484465873714311777937421161581798189554141474236239447612421883232173914183732126332838194648583472419154369952477422666389517569944428464617457124369349242479612422673241361777576466946622932243728551273284837934497511114334421486262244982914734452113946361245377351849815584855691778894798219822463298387771923329337634394654439458564233259451453345316753241438267739439225497515276522424441532462541528195782818326918562247278496495764435386667383577543385186827269732261223156824351164841648424564925198783625721396988984481558391866483955533972212164693898955412719161648411279149413443192896864258215498543827458438871355879336892721675937111952479183496982825163456282747678364612135596373533447719867384667516572262124225585623974278833981365494628646614588114147473559138853453189448624976774641922469183942857695986376428944876851497914443873513862319484181787593572987444669767939526294424531262999564948571142342741129862311311313166798363442745792896227642881893134498151552326647933689596516859342242244584714818773791567187322217164347852843751875979415198165627534263527828414549217234322361937785185174993256753483876378332521824515977173397535784236923629636713469151526399149548322849831431526219478653861754364155275865511643923249858589466142474763778413826829226663398467569555747267195129525138917561785436449855933951538973995881954521124753369223898312843734771532342383282987422334196585128526526324291777689689492346231786335851551413876834969878"
#count = 0
#for c in range(0, len(input)):
#    if input[c] == input[(c+1) % len(input)]:
#        count += int(input[c])
#print("Day 1-1: "+str(count))


#''' Day 1-2 '''
#count = 0
#half = len(input)//2
#for c in range(0, len(input)-1):
#    if input[c] == input[(c+half) % len(input)]:
#        count += int(input[c])
#print("Day 1-2: "+str(count))


#''' Day 2-1 '''
#input = "409,194,207,470,178,454,235,333,511,103,474,293,525,372,408,428\n4321,2786,6683,3921,265,262,6206,2207,5712,214,6750,2742,777,5297,3764,167\n3536,2675,1298,1069,175,145,706,2614,4067,4377,146,134,1930,3850,213,4151\n2169,1050,3705,2424,614,3253,222,3287,3340,2637,61,216,2894,247,3905,214\n99,797,80,683,789,92,736,318,103,153,749,631,626,367,110,805\n2922,1764,178,3420,3246,3456,73,2668,3518,1524,273,2237,228,1826,182,2312\n2304,2058,286,2258,1607,2492,2479,164,171,663,62,144,1195,116,2172,1839\n114,170,82,50,158,111,165,164,106,70,178,87,182,101,86,168\n121,110,51,122,92,146,13,53,34,112,44,160,56,93,82,98\n4682,642,397,5208,136,4766,180,1673,1263,4757,4680,141,4430,1098,188,1451\n158,712,1382,170,550,913,191,163,459,1197,1488,1337,900,1182,1018,337\n4232,236,3835,3847,3881,4180,4204,4030,220,1268,251,4739,246,3798,1885,3244\n169,1928,3305,167,194,3080,2164,192,3073,1848,426,2270,3572,3456,217,3269\n140,1005,2063,3048,3742,3361,117,93,2695,1529,120,3480,3061,150,3383,190\n489,732,57,75,61,797,266,593,324,475,733,737,113,68,267,141\n3858,202,1141,3458,2507,239,199,4400,3713,3980,4170,227,3968,1688,4352,4168"
#lines = input.split('\n')
#for i in range(len(lines)):
#    lines[i] = [int(j) for j in lines[i].split(',')]
#count = 0
#for line in lines:
#    count += max(line) - min(line)
#print("Day 2-1: "+str(count))


#''' Day 2-2 '''
#count = 0
#for line in lines:
#    for x in line:
#        possible = [z for z in line if z<x]
#        for y in possible:
#            if x%y == 0:
#                count += x/y
#print("Day 2-2: "+str(count))


#''' Day 3-1 '''
#input = 277678
#steps = 0
#square = 1
#while square**2 < input:
#    steps += 1
#    square += 2
#perim = square**2 - (square-2)**2
#extra = input - (square-2)**2
#length = perim//4
#dist_to_midpoint = abs( length/2 - (extra % length) )
#steps += dist_to_midpoint # straight line to the middle of the ring perimeter
#print("Day 3-1: "+str(steps))


#''' Day 3-2 '''
#input = 277678
#NORTH, S, W, E = (0, -1), (0, 1), (-1, 0), (1, 0) # directions
#turn_right = {NORTH: E, E: S, S: W, W: NORTH} # old -> new direction

#def spiral(width, height):
#    if width < 1 or height < 1:
#        raise ValueError
#    x, y = width // 2, height // 2 # start near the center
#    dx, dy = NORTH # initial direction

#    matrix = [[0] * width for _ in range(height)]
#    matrix[y][x] = 1 # initialize
#    while True:
#        surroundings = [matrix[y+i][x-1:x+2] for i in range(-1,2)] # 3x3 matrix of surrounding cells
#        matrix[y][x] = sum(sum(surroundings,[])) # non-zero = visited
#        #print(matrix[y][x])
#        if matrix[y][x] > input:
#            return matrix[y][x] # we can stop here for the puzzle. 
#        # try to turn right
#        new_dx, new_dy = turn_right[dx,dy]
#        new_x, new_y = x + new_dx, y + new_dy
#        if (0 <= new_x < width and 0 <= new_y < height and
#            matrix[new_y][new_x] == 0): # it's been visited; we can turn right
#            x, y = new_x, new_y
#            dx, dy = new_dx, new_dy
#        else: # try to move straight
#            x, y = x + dx, y + dy
#            if not (0 <= x < width and 0 <= y < height):
#                return matrix # nowhere to go

#print("Day 3-2: "+str(spiral(12,12))) # 12,12 hits everything below 8 million


#''' Day 4-1 '''
#count = 0
#with open('Day4.txt') as f:
#    data = f.read()
#    lines = data.split('\n')
#    for line in lines:
#        words = line.split(' ')
#        if len(words) == len(set(words)):
#            count += 1
#print("Day 4-1: "+str(count))


#''' Day 4-2 '''
#count = 0
#with open('Day4.txt') as f:
#    data = f.read()
#    lines = data.split('\n')
#    for line in lines:
#        words = line.split(' ')
#        letters = [''.join(sorted(word)) for word in words]
#        if len(letters) == len(set(letters)):
#            count += 1
#print("Day 4-2: "+str(count))


#''' Day 5-1 '''
#count = 0
#pos = 0
#with open('Day5.txt') as f:
#    data = f.read()
#    lines = data.split('\n')
#    jumps = [int(x) for x in lines]
#    while pos < len(jumps):
#        count += 1
#        jump = jumps[pos]
#        jumps[pos] += 1
#        pos += jump
#print("Day 5-1: "+str(count))


#''' Day 5-2 '''
#import time
#start = time.time()
#count = 0
#pos = 0
#with open('Day5.txt') as f:
#    data = f.read()
#    lines = data.split('\n')
#    jumps = [int(x) for x in lines]
#    while pos < len(jumps):
#        count += 1
#        jump = jumps[pos]
#        if jumps[pos] < 3:
#            jumps[pos] += 1
#        else:
#            jumps[pos] -= 1
#        pos += jump
#print("Day 5-2: "+str(count))
#print(time.time() - start)


#''' Day 6-1 '''
#import numpy as np
#banks = np.array([11,11,13,7,0,15,5,5,4,4,1,1,7,1,15,11])
#n = banks.size
#previous = []
#count = 0
#while not any( np.array_equal(banks,x) for x in previous):
#    previous.append(np.copy(banks))
#    count += 1
#    i = np.argmax(banks) # finds first occurance
#    blocks = banks[i]
#    banks[i] = 0
#    while blocks > 0:
#        i += 1
#        banks[i%n] += 1
#        blocks -= 1
#    print(banks)
#print("Day 6-1: "+str(count))

''' Day 6-2 '''
import numpy as np
banks = np.array([11,11,13,7,0,15,5,5,4,4,1,1,7,1,15,11])
n = banks.size
previous = []
while not any( np.array_equal(banks,x) for x in previous):
    previous.append(np.copy(banks))
    i = np.argmax(banks) # finds first occurance
    blocks = banks[i]
    banks[i] = 0
    while blocks > 0:
        i += 1
        banks[i%n] += 1
        blocks -= 1
    #print(banks)
matches = [np.array_equal(banks,x) for x in previous]
cycle = len(previous) - matches.index(True)
print("Day 6-2: "+str(cycle))