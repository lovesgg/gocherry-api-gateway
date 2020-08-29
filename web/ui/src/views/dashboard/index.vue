<template>
  <div class="dashboard-container">
    <markdown-preview :mktext="mktext" />
  </div>
</template>

<script>
import Axios from 'axios'
import { mapGetters } from 'vuex'
import { MarkdownPreview } from '@/components/index'

export default {
  name: 'Dashboard',
  components: {
    MarkdownPreview
  },
  data() {
    return {
      mktext: 'test'
    }
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  mounted() {
    this.getMdText()
  },
  methods: {
    getMdText() {
      const _this = this
      Axios.get('./help/readMe.md').then(res => {
        if (res.status === 200) {
          _this.mktext = res.data
        }
      })
    },
    createXhr() {
      try {
        return new XMLHttpRequest()
      } catch (e) {
        console.log(e)
      }
      // try {
      //   return new ActiveXObject('Microsoft.XMLHTTP')
      // } catch (e) {
      //   console.log(e)
      // }
      // alert("Ajax对象由于浏览器罕用而创建失败");
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~@/styles/defines.scss';
.dashboard {
  &-container {
    padding: 30px;
    width: 100%;
    height: 100%;
    background-color: $--my-background-color1;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
