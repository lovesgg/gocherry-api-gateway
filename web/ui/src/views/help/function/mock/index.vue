<template>
  <div class="help-function-mock">
    <markdown-preview :mktext="mktext" />
  </div>
</template>

<script>
import Axios from 'axios'
import { mapGetters } from 'vuex'
import { MarkdownPreview } from '@/components/index'

export default {
  name: 'HelpProfileMock',
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
      Axios.get('./help/function/mock/index.md').then(res => {
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
.help-function-mock {
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
