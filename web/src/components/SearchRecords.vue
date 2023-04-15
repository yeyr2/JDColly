<template>
  <div>
    <div class="main">
      <div class="popularTags">
        <h1>热门标签</h1>
        <div class="tags">
          <el-tag
            type="info"
            v-for="(item, index) in tagsArr"
            :key="index"
            @click="search(item)"
          >
            {{ item }}
          </el-tag>
        </div>
      </div>
      <div class="historyRecords">
        <h1>搜索记录</h1>
        <div class="records">
          <table border="1" cellspacing="0">
            <tr>
              <th>搜索内容</th>
              <th>搜索时间</th>
            </tr>
            <tr v-for="(item, index) in recordsArr" :key="index">
              <td>{{ item.key }}</td>
              <td>{{ item.time }}</td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      tagsArr: ["iPhone", "iPad", "Watch", "AirPods"],
      recordsArr: [],
    };
  },
  async created() {
    const { data: data } = await this.$http.get(`new/search`, {
      params: {
        id: localStorage.getItem("id"),
        token: localStorage.getItem("token"),
      },
    });
    this.recordsArr = data.value;
  },
  methods: {
    search(item) {},
  },
};
</script>
<style scoped>
.main {
  width: 50%;
  margin: 0 auto;
}
/* 热门标签 */
.main .popularTags {
  width: 100%;
  /* height: 200px; */
  /* background-color: aqua; */
  padding: 20px 50px;
}
.main .popularTags h1 {
  margin-bottom: 20px;
}
.main .popularTags .tags {
  width: 100%;
  height: 30px;
  padding-left: 20px;
  /* background-color: aqua; */
}
.el-tag {
  height: 30px;
  margin-right: 30px;
  line-height: 30px;
}
/* 历史记录 */
.main .historyRecords {
  width: 100%;
}
.main .historyRecords h1 {
  padding: 0 50px 20px;
}
.main .historyRecords .records {
  width: 100%;
  padding-left: 20px;
}
tr:hover {
  background-color: #eeeeee;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  text-align: center;
  font-size: 18px;
  font-weight: normal;
}
table td,
table th {
  border: 1px solid #cad9ea;
  color: #666;
  height: 30px;
}
table thead th {
  background-color: #cce8eb;
  width: 100px;
}
table tr:nth-child(odd) {
  background: #fff;
}
table tr:nth-child(even) {
  background: #f5fafa;
}
</style>