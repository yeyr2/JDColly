#!/usr/bin/python
# -*- coding: utf-8 -*-
# coding=utf-8
import os
import platform
from collections import Counter

import jieba
from wordcloud import WordCloud

font_path_ubuntu = "usr/share/fonts/truetype/wqy/wqy-zenhei.ttc"
font_path_arch = "/usr/share/fonts/wenquanyi/wqy-zenhei/wqy-zenhei.ttc"


def split_text(text):
    # split_text函数用于jieba分词。
    text = list(text)
    words = ""

    for element in text:
        tmp = jieba.cut_for_search(element)
        # 用Counter方法计算单词频率数
        count = Counter(tmp)
        most_count = count.most_common()

        for i in most_count:
            if i[0] == '\n':
                continue
            words = "\n".join([words, i[0]])

    return words


def word_cloud(text_str, productId):
    if platform.platform().find("arch") != -1:  # arch
        font_path = font_path_arch
    else:
        font_path = font_path_ubuntu  # ubuntu

    stopwords = set(map(str.strip, open('./stopwords/cn_stopwords.txt').readlines(),
                        open("./stopwords/baidu_stopwords.txt").readlines()))
    wc = WordCloud(background_color="#0C1853",  # 设置背景颜色
                   max_words=2000,  # 词云显示的最大词数
                   height=400,  # 图片高度
                   width=400,  # 图片宽度
                   max_font_size=50,  # 最大字体
                   stopwords=stopwords,  # 设置停用词
                   font_path=font_path,  # 兼容中文字体，不然中文会显示乱码
                   ).generate(text_str)  # 此处的text便是分好词的19大文本

    # 生成的词云图像保存到本地
    folder_name = "./images"
    if not os.path.exists(folder_name):
        os.makedirs(folder_name)
        print(f"Folder {folder_name} has been created.")
    else:
        print(f"Folder {folder_name} already exists.")
    paths = "img" + productId + ".jpg"
    wc.to_file("./images/" + paths)

    # 显示图像
    # plt.imshow(wc, interpolation='bilinear')
    # interpolation ='bilinear' # 表示插值方法为双线性插值
    # plt.axis("off")  # 关掉图像的坐标
    # plt.show()
    return paths


if __name__ == '__main__':
    # file = open('./wenzi/tt2', encoding='utf8').read()
    file = ["早上好", "下午好下午打", "晚上好"]
    string = split_text(text=file)
    path = word_cloud(string, "12")
