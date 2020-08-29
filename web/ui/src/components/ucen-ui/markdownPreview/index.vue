<template>
  <div class="ucen-preview" v-html="mkhtml" />
</template>
<script type="text/javascript">
import './index.scss'
// import hljs from 'pluginspath/highlight/highlight.pack.js'
var marked = require('marked')
marked.setOptions({
  renderer: new marked.Renderer(),
  pedantic: false,
  gfm: true,
  tables: true,
  breaks: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  xhtml: false
  // highlight: function (code) {
  //   return hljs.highlightAuto(code).value;
  // }
})
export default {
  name: 'UcenPreview',
  props: {
    mktext: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      mkhtml: ''
    }
  },
  watch: {
    // 'mktext' : 'mktextChange'
    mktext: {
      handler(val, oval) {
        this.mktextChange()
      },
      immediate: true,
      deep: true
    }
  },
  mounted() {
  },
  methods: {
    mktextChange() {
      if (this.mktext) {
        this.mkhtml = marked(this.mktext)
      } else {
        this.mkhtml = ''
      }
    }
  }
}
</script>
<style lang="scss" scoped>
@import "index.scss";
</style>
