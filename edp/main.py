from fipy import *
from math import sqrt
import matplotlib.tri as mtri
import matplotlib.pyplot as plt
import numpy as np
import time

grid_size=0.01
D = 0.2 # 20 cm
R = D/2
mu = 80e9 # G típico de acero


def run_simulation(k):
    hexagon_script = f'''
    // Gmsh geometry script for a regular hexagon
    R = {R};

    Point(1) = {{ R,           0, 0, 0.1 }};
    Point(2) = {{ R/2,  R*Sqrt(3)/2, 0, 0.1 }};
    Point(3) = {{-R/2,  R*Sqrt(3)/2, 0, 0.1 }};
    Point(4) = {{-R,           0, 0, 0.1 }};
    Point(5) = {{-R/2, -R*Sqrt(3)/2, 0, 0.1 }};
    Point(6) = {{ R/2, -R*Sqrt(3)/2, 0, 0.1 }};

    Line(1) = {{1, 2}};
    Line(2) = {{2, 3}};
    Line(3) = {{3, 4}};
    Line(4) = {{4, 5}};
    Line(5) = {{5, 6}};
    Line(6) = {{6, 1}};

    Line Loop(1) = {{1, 2, 3, 4, 5, 6}};
    Plane Surface(1) = {{1}};

    Mesh.CharacteristicLengthFactor = 0.2;
    Mesh.MeshSizeMin = {grid_size/2};
    Mesh.MeshSizeMax = {grid_size};
    Mesh 2;
    '''
    mesh = Gmsh2D(hexagon_script)
    phi = CellVariable(name = "phi", mesh = mesh, value = 0.5)
    eq = (-2*mu*k == DiffusionTerm(coeff=1., var=phi))

    phi.constrain(0., where=mesh.exteriorFaces)

    for i in range(10):
        eq.solve(var=phi)

    T13 = phi.grad[1]
    T23 = -phi.grad[0]

    viewer = Viewer(vars=phi)
    viewer.plot(filename=f"plots/k{k}phi.png")
    viewer13 = Viewer(vars=T13)
    viewer13.plot(filename=f"plots/k{k}T13.png")
    viewer23 = Viewer(vars=T23)
    viewer23.plot(filename=f"plots/k{k}T23.png")

    T = 2 * (phi * mesh.cellVolumes).sum()
    print(f"{k}\t{T}")


print("k\tT")
run_simulation(0.01)
run_simulation(0.1)
run_simulation(0.5)
run_simulation(1)
run_simulation(5)
run_simulation(10)
run_simulation(20)
run_simulation(50)
run_simulation(100)
run_simulation(500)
