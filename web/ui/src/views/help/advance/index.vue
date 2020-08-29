<template>
  <div style="padding:30px;" class="help-advance">
    <el-alert :closable="false" :title="title">
      <markdown-preview :mktext="mktext" style="margin-bottom: 20px;" />
      <router-view />
    </el-alert>
  </div>
</template>

<script>
import Axios from 'axios'
import { mapGetters } from 'vuex'
import { MarkdownPreview } from '@/components/index'

export default {
  components: {
    MarkdownPreview
  },
  data() {
    return {
      title: '无',
      mktext: ''
    }
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  mounted() {
    this.getRouterTitle()
    this.getMdText()
  },
  methods: {
    getRouterTitle() {
      const matched = this.$route.matched.filter(item => item.meta && item.meta.title)
      const levelList = matched.filter(item => item.meta && item.meta.title && item.meta.breadcrumb !== false)
      if (levelList && levelList.length > 0) {
        this.title = levelList[levelList.length - 1].meta.title
      }
    },
    getMdText() {
      const _this = this
      Axios.get('./help/advance/index.md').then(res => {
        if (res.status === 200) {
          _this.mktext = res.data
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~@/styles/defines.scss';
.help-advance {
  .el-alert {
    background-color: $--my-background-color1;
  }
}
</style>
