const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    // proxy: {
    //   '/api': {
    //     target: 'http://47.100.16.28:9090/api',
    //     changeOrigin: true,
    //     pathRewrite: {
    //       '^api': ''
    //     }
    //   }
    // }
  }
})

