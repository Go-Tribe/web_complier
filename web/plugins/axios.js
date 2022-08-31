export default function({$axios}, inject) {
  const $http = $axios.create({
    timeout: 3000,
    withCredentials: true,
  });

  $http.onRequest((config) => {
    console.log('config', config);
  });

  $http.onResponse((response) => {
    console.log('response', response);
    let err = null;
    let errMsg = '';
    if(response.status === 200) {
      // http请求正常返回
      const {code, data, message} = response.data;
      switch(code) {
        case 0: {
          return data;
        }
        default: {
          errMsg = code;
          err = new Error(errMsg);
          break;
        }
      }
    } else {
      switch(response.status) {
        case 403: {
          errMsg = '403-无权限访问';
          err = new Error(errMsg);
          break;
        }
        case 404: {
          errMsg = '404-not found';
          err = new Error(errMsg);
          break;
        }
        default: {
          errMsg = `${response.status}-服务异常`;
          err = new Error(errorMsg);
          break;
        }
      }
    }
    return Promise.reject(err);
  });

  $http.onResponseError((error) => {
    if(error && error.response) {
      if(error.response.status !== 401) {
        console.error('response error occurred: %d--%s', error.response.status, error);
      }
    } else if(error && error.code) {
      switch(error.code) {
        case 'ECONNABORTED': {
          console.error('服务异常，请重试');
          break;
        }
        default: {
          // do nothing
          break;
        }
      }
    } else {
      // nothing
    }
  });

  $http.onError((error) => {
    console.log(error);
  })

  inject('http', $http);
}