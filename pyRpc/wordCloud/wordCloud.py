#!/usr/bin/python
# -*- coding: utf-8 -*-
# coding=utf-8

import platform
from collections import Counter
import jieba
from wordcloud import WordCloud

font_path_ubuntu = "usr/share/fonts/truetype/wqy/wqy-zenhei.ttc"
font_path_arch = "/usr/share/fonts/wenquanyi/wqy-zenhei/wqy-zenhei.ttc"


def split_four_text(text):
    # split_four_text函数用于jieba分词并分隔为4个字为一组的内容。

    words = jieba.cut(text)

    # 用Counter方法计算单词频率数
    count = Counter(words)
    most_count = count.most_common()
    words_list = []

    for i in most_count:
        if len(i[0]) == 4:
            words_list.append(i[0])

    return words_list


def word_cloud(text):

    if platform.platform().find("arch") != -1:  # arch
        font_path = font_path_arch
    else:
        font_path = font_path_ubuntu  # ubuntu

    stopwords = set(map(str.strip, open('../stopwords/cn_stopwords.txt').readlines()))
    wc = WordCloud(background_color="white",  # 设置背景颜色
                   max_words=2000,  # 词云显示的最大词数
                   height=400,  # 图片高度
                   width=800,  # 图片宽度
                   max_font_size=50,  # 最大字体
                   stopwords=stopwords,  # 设置停用词
                   font_path=font_path,  # 兼容中文字体，不然中文会显示乱码
                   ).generate(text)  # 此处的text便是分好词的19大文本

    # 生成的词云图像保存到本地
    wc.to_file("../images/1.jpg")

    # 显示图像
    # plt.imshow(wc, interpolation='bilinear')
    # interpolation ='bilinear' # 表示插值方法为双线性插值
    # plt.axis("off")  # 关掉图像的坐标
    # plt.show()


if __name__ == '__main__':
    text = open('../wenzi/tt2', encoding='utf8').read()
    word_cloud(str(split_four_text(text=text)))
