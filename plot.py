import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import pandas as pd

# ファイル名を標準入力
sch = input()
filepath = './'+sch+'result.csv'

df = pd.read_csv(filepath ,header = None)
df = df.T # 座標と時間逆だったので転置を取る
Nx = len(df.index)
Ny = len(df.columns)
x=list(range(Nx))
y=list(range(Ny))
X, Y = np.meshgrid(x,y)
u = df.values

def functz(u):
    z=u[X,Y]
    return z

Z = functz(u)
fig = plt.figure()
ax = Axes3D(fig)
ax.plot_wireframe(X,Y,Z, color='r')
ax.set_xlabel('x')
ax.set_ylabel('t')
ax.set_zlabel('U')

plt.show()
