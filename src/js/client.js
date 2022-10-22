'use strict';

import axios from 'axios';

import AppError from './error.js';

class AppClient {
  constructor() {
    this.axios = axios.create();
    this.axios.interceptors.response.use((response) => {
      let data = response.data;
      if (data.state === 'error') {
        let result = data.result;
        return Promise.reject(new AppError(result.message, result.code));
      }
      return data.result;
    });
  }

  initToken() {
    delete this.axios.defaults.headers.common['Authorization'];
  }

  setToken(token) {
    this.axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
  }

  challenge(address) {
    let p = new URLSearchParams();
    p.append('user_address', address);
    console.log("CHALLENING",p);

    return this.axios.post('/auth/challenge', p);
  }

  authorize(address, signature) {
    let p = new URLSearchParams();
    p.append('user_address', address);
    p.append('user_signature', signature);
    console.log("AUTHING",p);

    return this.axios.post('/auth/authorize', p);
  }

  getUser(address) {
    return this.axios.get('/api/users/' + address);
  }

  updateUser(address, params) {
    let p = new URLSearchParams();
    p.append('user_first_name', params.userFirstName);
    p.append('user_last_name', params.userLastName);
    p.append('message_to_address', params.messageToAddress);
    p.append('message_subject', params.messageSubject);
    p.append('message_body', params.messageBody);
    p.append('user_address', params.userAddress);
    console.log("UPDATING",p);
    return this.axios.put('/api/users/' + address, p);
  }

  deleteUser(address) {
    return this.axios.delete('/api/users/' + address);
  }
}

export default AppClient;
