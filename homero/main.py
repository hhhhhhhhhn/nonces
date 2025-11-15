from collections import deque

def closest_paths(n,k):
    paths = [-1 for i in range(n+1)]             # Inicializar caminos más cortos
    paths[0] = 0                                 # (Partimos asumiendo inaccesibles (-1), excepto el nodo inicial)

    to_visit = deque([0])                        # Inicializar cola de nodos por recorrer

    while len(to_visit) > 0:                     # Mientras queden nodos por recorrer
        Q = to_visit.popleft()                   #   Visitar el primero en la cola
        for i in range(0, k+1):                  #   Por cada posible arista que tiene
            if k - Q <= i and i <= n - Q:        #      Si esa arista efectivamente existe
                new_Q = Q - k + 2*i              #      Calcular a el cual la arista conecta
                if paths[new_Q] == -1:           #      Y si no ha sido visitado
                    paths[new_Q] = Q             #        Su camino más rápido es a través del nodo actual
                    to_visit.append(new_Q)       #        Y ponemos el nuevo nodo en el fondo de la cola por visitar

    return paths

def homero(n, k, switches):
    paths = closest_paths(n, k)                  # Cálculo de caminos más cortos
    Q = sum([1 for i in switches if i == "ON"])  # Cálculo de vértice actual en grafo
    if paths[Q] == -1:                           # Si no existe camino, retornamos
        return "ERROR: No es posible"

    sequence = []                                # Inicialización de secuencia de pasos

    while Q != 0:                                                      # Mientras haya interruptores prendidos
        new_Q = paths[Q]                                               # Buscamos a qué vértice tenemos que ir despues,
        amount_to_turn_off = (Q - new_Q + k)//2                        # Y cuántos interruptores hay que apagar y prender.
        amount_to_turn_on = k - amount_to_turn_off
        switches_to_switch = []                                        # Iniciamos la lista de interruptores

        for i in range(len(switches)):                                 # Iteramos por los elementos de la lista
            if switches[i] == "OFF" and amount_to_turn_on > 0:         #  Si está en OFF y nos falta prender, lo prendemos
                switches[i] = "ON"
                amount_to_turn_on -= 1
                switches_to_switch.append(i+1)                         #  (El +1 es para que los índices partan en 1)
            elif switches[i] == "ON" and amount_to_turn_off > 0:       #  Si está en ON y nos falta apagar, lo apagamos
                switches[i] = "OFF"
                amount_to_turn_off -= 1
                switches_to_switch.append(i+1)
            if amount_to_turn_on == 0 and amount_to_turn_off == 0:     #  Y ya no hay por prender y apagar, terminamos
                break
        sequence.append(switches_to_switch)                            # Añadimos los interruptores que cambiamos a la secuencia de pasos
        Q = new_Q                                                      # Y repetimos con el nuevo número de interruptores prendidos
    return sequence

assert len(homero(5, 2, ["ON", "OFF", "ON", "OFF", "OFF"])) == 1
assert len(homero(5, 3, ["ON", "OFF", "ON", "OFF", "OFF"])) == 2
