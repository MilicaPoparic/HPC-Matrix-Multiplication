import math, numpy as np
import matplotlib.pyplot as plt

def calc_amdal(cpu):
    amdal = []
    for c in cpu:
        amdal.append(1 / (0.01 + 0.99 / c))
    return amdal

def plot_graph(speedUp, law, cpu, type):
    plt.figure(0)
    plt.plot(cpu, speedUp, color='blue', marker='o')
    plt.plot(cpu, law, color='red', marker='o', linestyle='dashed')

    plt.title(type + " skaliranje", {'size': '14'})
    plt.grid('on')
    plt.axis('equal')
    plt.xlabel("Broj jezgara", {'size': '14'})
    plt.ylabel("Ubrzanje", {'size': '14'})
    plt.savefig('resources/strong_scaling.png', bbox_inches='tight')
    plt.show()

def calculate(fileParalel, fileSeq):
    sumS = 0
    sumP = 0
    listaS = []
    listaP = []
    f = open(fileParalel, "r")
    for line in f:
        sumP += float(line.strip())
        listaP.append(float(line.strip()))
    f.close()
    avgP = sumP/30
    print(avgP, "Mean parallel")
    print(np.std(listaP), "Std parallel")

    f1 = open(fileSeq, "r")
    for line in f1:
        sumS += float(line.strip())
        listaS.append(float(line.strip()))
    f1.close()
    avgS = sumS / 30
    print(avgS, "Mean sequential")
    print(np.std(listaS), "std sequential")

    print(avgS/avgP, "SpeedUp")

    return avgS/avgP