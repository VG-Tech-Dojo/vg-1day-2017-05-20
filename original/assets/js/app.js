(function() {
  'use strict';
  const Message = function() {
    this.body = '';
  };

  Vue.component('message', {
    // 1-1. ユーザー名を表示しよう
	// 1-2. ユーザー名を追加しよう
    props: ['id', 'body', 'removeMessage'],
    template: `
    <div class="message">
      <span>{{ body }}</span>
      <span class="remove-message-button u-pull-right" v-on:click="remove">x</span>
    </div>
  `,
    methods: {
      remove() {
        this.removeMessage(this.id);
      }
    }
  });

  const app = new Vue({
    el: '#app',
    data: {
      messages: [],
      newMessage: new Message()
    },
    created() {
      this.getMessages();
    },
    methods: {
      getMessages() {
        fetch('/api/messages').then(response => response.json()).then(data => {
          this.messages = data.result;
        });
      },
      sendMessage() {
        const message = this.newMessage;
        fetch('/api/messages', {
          method: 'POST',
          body: JSON.stringify(message)
        })
          .then(response => response.json())
          .then(response => {
            if (response.error) {
              alert(response.error.message);
              return;
            }
            this.messages.push(response.result);
            this.clearMessage();
          })
          .catch(error => {
            console.log(error);
          });
      },
      removeMessage(id) {
        fetch(`/api/messages/${id}`, {
          method: 'DELETE'
        })
          .then(response => {
            // TODO: 削除処理書く
            console.log(response);
          })
          .catch(error => {
            console.log(error);
          });
      },
      clearMessage() {
        this.newMessage = new Message();
      }
	  // 1-3. メッセージを編集しよう
      // ...
    }
  });
})();
