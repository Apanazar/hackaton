import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns


with open('db.json', encoding='utf-8') as inputfile:
    df = pd.read_json(inputfile)
df.to_csv('test.csv', encoding='utf-8', index=False)

news_d = pd.read_csv("train.csv")

sns.countplot(x="label", data=news_d)
print("1: Unreliable")
print("0: Reliable")
print(news_d.label.value_counts())