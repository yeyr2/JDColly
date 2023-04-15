<template>
  <div>
    <div class="top">
      <div class="content">商品评论评分可视化</div>
      <div class="time">
        当前时间： {{ this.year }}年{{ this.month }}月{{ this.day }}日
      </div>
    </div>
    <div class="main">
      <el-row :gutter="25" style="margin-bottom: 10px">
        <el-col :span="12">
          <div class="box-pie" style="height: 500px" ref="chart1"></div>
        </el-col>
        <el-col :span="12">
          <div class="box-pie" style="height: 500px" ref="chart3"></div
        ></el-col>
      </el-row>
      <el-row :gutter="25">
        <el-col :span="12">
          <div class="box-pie" style="height: 500px" ref="chart2"></div>
        </el-col>
        <el-col :span="12">
          <div class="box-pie" style="height: 500px" ref="chart4"></div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "",
  data() {
    return {
      chart1: null,
      chart3: null,
      chart2: null,
      year: "",
      month: "",
      day: "",
      info: null,
    };
  },
  computed: {
    // 分段分数折线图
    chartOption1() {
      return {
        xAxis: {
          type: "category",
          data: ["0-20", "20-40", "40-60", "60-80", "80-100"],
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            data: this.info.interval,
            type: "bar",
            showBackground: true,
            backgroundStyle: {
              color: "rgba(180, 180, 180, 0.2)",
            },
          },
        ],
      };
    },
    chartOption2() {
      return {
        tooltip: {
          trigger: "item",
        },
        legend: {
          top: "5%",
          left: "center",
        },
        series: [
          {
            name: "Battery Life",
            type: "pie",
            radius: ["40%", "70%"],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: "#fff",
              borderWidth: 2,
            },
            label: {
              show: false,
              position: "center",
            },
            emphasis: {
              label: {
                show: true,
                fontSize: "30",
                fontWeight: "bold",
                color: "auto",
              },
            },
            labelLine: {
              show: false,
            },
            data: [
              { value: this.info.interval[0], name: "0-20" },
              { value: this.info.interval[1], name: "20-40" },
              { value: this.info.interval[2], name: "40-60" },
              { value: this.info.interval[3], name: "60-80" },
              { value: this.info.interval[4], name: "80-100" },
            ],
          },
        ],
      };
    },
    // 总分数图表
    chartOption3() {
      return {
        series: [
          {
            type: "gauge",
            startAngle: 180,
            endAngle: 0,
            center: ["50%", "75%"],
            radius: "90%",
            min: 0,
            max: 100,
            splitNumber: 8,
            axisLine: {
              lineStyle: {
                width: 6,
                color: [
                  [0.25, "#FF6E76"],
                  [0.5, "#FDDD60"],
                  [0.75, "#58D9F9"],
                  [1.0, "#7CFFB2"],
                ],
              },
            },
            pointer: {
              icon: "path://M12.8,0.7l12,40.1H0.7L12.8,0.7z",
              length: "12%",
              width: 20,
              offsetCenter: [0, "-60%"],
              itemStyle: {
                color: "inherit",
              },
            },
            axisTick: {
              length: 12,
              lineStyle: {
                color: "inherit",
                width: 2,
              },
            },
            splitLine: {
              length: 20,
              lineStyle: {
                color: "inherit",
                width: 5,
              },
            },
            axisLabel: {
              // color: "#464646",
              color: "#ffffff",
              fontSize: 20,
              distance: -60,
              rotate: "tangential",
              formatter: function (value) {
                if (value === 100) {
                  return "好 评 如 潮";
                } else if (value === 75) {
                  return "商 品 不 错";
                } else if (value === 50) {
                  return "中 规 中 矩";
                } else if (value === 25) {
                  return "商 品 堪 忧";
                } else if (value === 0) {
                  return "差 评 预 警";
                }
                return "";
              },
            },
            title: {
              offsetCenter: [0, "-10%"],
              fontSize: 20,
            },
            detail: {
              fontSize: 30,
              offsetCenter: [0, "-35%"],
              valueAnimation: true,
              formatter: function (value) {
                return Math.round(value) + "";
              },
              color: "inherit",
            },
            data: [
              {
                value: parseInt(this.info.fraction),
                // value: 30,
                name: "Comments Rating",
              },
            ],
          },
        ],
      };
    },
    chartOption4() {
      return {
        series: [
          {
            type: "gauge",
            axisLine: {
              lineStyle: {
                width: 30,
                color: [
                  [0.3, "#a0a7e6"],
                  [0.7, "#3fb1e3"],
                  [1, "#6be6c1"],
                ],
              },
            },
            pointer: {
              itemStyle: {
                color: "auto",
              },
            },
            axisTick: {
              distance: -30,
              length: 8,
              lineStyle: {
                color: "#fff",
                width: 2,
              },
            },
            splitLine: {
              distance: -30,
              length: 30,
              lineStyle: {
                color: "#fff",
                width: 4,
              },
            },
            axisLabel: {
              color: "auto",
              distance: 35,
            },
            detail: {
              valueAnimation: true,
              formatter: "{value}",
              color: "auto",
              fontSize: 24,
              offsetCenter: [0, "80%"],
            },
            data: [
              {
                value: parseInt(this.info.fraction),
              },
            ],
          },
        ],
      };
    },
  },
  mounted() {
    this.chart1 = this.$echarts.init(this.$refs.chart1, "dark");
    this.chart2 = this.$echarts.init(this.$refs.chart2, "dark");
    this.chart3 = this.$echarts.init(this.$refs.chart3, "dark");
    this.chart4 = this.$echarts.init(this.$refs.chart4, "dark");
    this.chart1.setOption(this.chartOption1);
    this.chart2.setOption(this.chartOption2);
    this.chart3.setOption(this.chartOption3);
    this.chart4.setOption(this.chartOption4);
  },
  created() {
    this.info = JSON.parse(window.localStorage.getItem("shopComments"));
    const time = new Date();
    this.year = time.getFullYear();
    this.month = time.getMonth() + 1;
    this.day = time.getDate();
  },
  // methods: {},
};
</script>
<style scoped>
.top {
  width: 100%;
  height: 50px;
  text-align: center;
  line-height: 50px;
  background-color: #0c1853;
  color: #fff;
  position: relative;
  overflow: hidden;
}
.top .content {
  float: left;
  width: 100%;
  font-size: 20px;
  margin: 0 auto;
}
.top .el-card {
  float: left;
}
.top .time {
  position: absolute;
  float: right;
  right: 100px;
  font-size: 12px;
}
.main {
  width: 97%;
  margin: 20px auto 0;
  box-sizing: border-box;
}
</style>