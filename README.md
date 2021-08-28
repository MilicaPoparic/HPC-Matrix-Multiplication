# Kanonov algoritam množenja matrica

## Opis algoritma

Kanonov algoritam za množenje matrica izvršava se u n koraka gde je n broj dimenzija ulaznih(kvadratnih) matrica. U svakom koraku ulazne matrice se šiftuju na sledeći način: svaki red matrice A šiftuje se u levo, dok se svaka kolona matrice B šiftuje na gore, zatim se dobijene šiftovane matrice množe na sledeći način: C[i][j] = A[i][j] * B[i][j], tako da u svakom koraku dobijemo matricu C kao rezultat množenja. Krajnji rezultat algoritma dobija se sabiranjem svih C matrica.
Inicijalni korak se razlikuje od ostalih po načinu šiftovanja, u inicijalnom odnosno prvom koraku svaki od redova matrice A šiftuje se u levo za redni broj tog reda (0. red 0 shift left, 1. red 1 shift left...), slično je i sa matricom B (0. kolona 0 shift up, 1. kolona 1 shift up...), ostalih n-1 koraka imaju fiksni šift, odnosno svaki red/kolona šiftuju se za 1 poziciju levo/gore. 

## Sekvencijalna verzija

Sekvencijalna verzija algoritma izvršava iterativno svaki od koraka algoritma.

## Paralelna verzija
 
Paralelna verzija realizovana je upotrebom _MPI for Python_ paketa. Izvršavanje algoritma podeljeno je na p + 1 procesa. Svaki proces šiftuje i množi blokove matrica A i B, veličine n/sqrt(p) i kao rezultat daje maticu C. Rezultat algoritma dobija se na sličan način, svaka podmatrica C koja je rezultat množenja bloka sabira se sa odgovarajućim blokovima iz ostalih procesa, na kraju se C blokovi sklope u jednu rezultujuću matricu.



