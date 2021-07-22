<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <div slot="header">项目</div>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                
              </el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="8">
                
              </el-col>
              <el-col :span="8">
                
              </el-col>
              <el-col :span="8">
                
              </el-col>
            </el-row>
          </div>
        </el-card>
        <el-card style="margin-top: 20px">
          <div slot="header">团队</div>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                
              </el-col>
            </el-row>
            <el-row style="margin-left: 40px" :gutter="20">
              <template v-for="(item, index) in members">
                <el-col :key="index" :span="8">
                  <a :href="item.html_url">
                    <img class="avatar-img" :src="item.avatar_url">
                    <a class="author-name" style="">{{ item.login }}</a>
                  </a>
                </el-col>
              </template>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <div slot="header">
            提交记录
          </div>
          <div>
            <Timeline
              :timeline-items="dataTimeline"
              :message-when-no-items="messageWhenNoItems"
              :unique-timeline="true"
              :unique-year="true"
              :show-day-and-month="true"
              order="desc"
              date-locale="zh-CN"
            />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Timeline from 'timeline-vuejs'
export default {
  name: 'About',
  components: {
    Timeline
  },
  data() {
    return {
      messageWhenNoItems: 'There arent commits',
      members: [],
      dataTimeline: [],
      page: 0
    }
  },
  created() {

  },
  methods: {

  }
}
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
