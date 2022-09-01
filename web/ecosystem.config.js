module.exports = {
apps: [
    {
      name: 'complier-web',
      exec_mode: 'cluster',
      instances: 'max', // Or a number of instances
      script: './node_modules/nuxt/bin/nuxt.js',
      args: 'start',
      watch: true, // 启动热重载
      // 配置环境变量，这里的环境变量要与nuxt里边的`package.json`文件的变量相同
      error_file: '../logs/complier-web-err.log',
      out_file: '../logs/complier-web-out.log',
      pid_file: '../pids/pid.log',
      min_uptime: '200s',
      ignore_watch: ['node_modules', 'logs'],
      max_restarts: 10,
      max_memory_restart: '300M',
      merge_logs: true,
      log_date_format: 'YYYY-MM-DD HH:mm Z',
    },
  ]
}