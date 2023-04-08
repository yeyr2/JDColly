from snownlp import SnowNLP

import service.analyze_pb2 as analyze_pb2


def analysis_comment(request):
    # 进行情感分析
    analyze_comment = analyze_pb2.AnalyzeComment()

    count = [0, 0, 0, 0, 0]
    total = [0, 0, 0, 0, 0]
    lists = list(request)
    for text in lists:
        s = SnowNLP(text).sentiments
        if s < 0.2:
            total[0] += s
            count[0] += 1
        elif s < 0.4:
            total[1] += s
            count[1] += 1
        elif s < 0.6:
            total[2] += s
            count[2] += 1
        elif s < 0.8:
            total[3] += s
            count[3] += 1
        else:
            total[4] += s
            count[4] += 1

    interval = []
    totals = int(0)
    counts = int(0)
    for i in range(5):
        total[i] = int(total[i] * 100)
        totals += total[i]
        counts += int(count[i])
        if count[i] == 0:
            interval.append(0)
            continue
        interval.append(count[i])

    if counts == 0:
        analyze_comment.Fraction = 0
    else:
        analyze_comment.Fraction = int(totals) // int(counts)
    analyze_comment.Interval.extend(interval)
    return analyze_comment


# if __name__ == '__main__':
#     r = analysis_comment(["早上好","晚上好"])
#     print(r.Fraction,r.Interval)
