<template>
  <div class="help-function-envvar">
    <markdown-preview :mktext="mktext" />
  </div>
</template>

<script>
import Axios from 'axios'
import { mapGetters } from 'vuex'
import { MarkdownPreview } from '@/components/index'

export default {
  name: 'HelpProfileEnvvar',
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
      Axios.get('./help/function/envVar/index.md').then(res => {
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
.help-function-envvar {
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
