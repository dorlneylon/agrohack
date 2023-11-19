import json
import numpy as np
from scipy.interpolate import LinearNDInterpolator
import os


class WeatherMatrixBuilder:
    def __init__(self, file_prefix, file="stations.json") -> None:
        self.file_pref = file_prefix
        with open(file_prefix + file) as f:
            self.s = json.load(f)
        with open(file_prefix + "cfg.json") as f:
            cfg = json.load(f)
        latitudes = []
        longitudes = []
        self.indexes = []
        for i in self.s.keys():
            self.indexes.append(i)
            latitudes.append(self.s[i][1])
            longitudes.append(self.s[i][2])
        self.maxlatitude = max(latitudes)
        self.maxlongitude = max(longitudes)
        self.minlatitude = min(latitudes)
        self.minlongitude = min(longitudes)
        self.gr1 = np.linspace(self.minlatitude, self.maxlatitude, num=cfg["width"])
        self.gr2 = np.linspace(self.minlongitude, self.maxlongitude, num=cfg["height"])
        self.x1, self.y1 = np.meshgrid(self.gr1, self.gr2)

    def produce(self, time, data="new_data"):
        data = self.file_pref + data
        t = []
        u = []
        latitudes = []
        longitudes = []
        for i in self.indexes:
            path = f"{data}/{i}--{time}.json"
            if os.path.exists(path):
              with open(path) as f:
                j = json.load(f)
              latitudes.append(j[3])
              longitudes.append(j[4])
              t.append(j[6]["T"])
              u.append(j[6]["U"])
        t_interpolator = LinearNDInterpolator(list(zip(latitudes, longitudes)), t)
        u_interpolator = LinearNDInterpolator(list(zip(latitudes, longitudes)), u)

        z_t = t_interpolator(self.x1, self.y1)
        z_u = u_interpolator(self.x1, self.y1)
        np.save(self.file_pref + f"days_log/{time}.npy", np.array([z_t, z_u]))