(function() {
  'use strict';
  const Message = function() {
    this.body = '';
  };

  Vue.component('message', {
    // 1-1. ユーザー名を表示しよう
    props: ['id', 'body', 'messageSender', 'removeMessage', 'updateMessage'],
    data() {
      return {
        editing: false,
        editedBody: null,
        displayedBody: this.body,
        messageSender: this.messageSender
      }
    },
    // 1-1. ユーザー名を表示しよう
    template: `
    <div class="message">
      <div v-if="editing">
        <div class="row">
          <textarea v-model="editedBody" class="u-full-width"></textarea>
          <button v-on:click="doneEdit">Save</button>
          <button v-on:click="cancelEdit">Cancel</button>
        </div>
      </div>
      <div class="message-body" v-else>
        <span>{{ displayedBody }}</span>
        <span>{{ messageSender }}</span>
        <span class="action-button u-pull-right" v-on:click="edit">&#9998;</span>
        <span class="action-button u-pull-right" v-on:click="remove">&#10007;</span>
      </div>
    </div>
  `,
    methods: {
      remove() {
        this.removeMessage(this.id)
          .then(() => {
            console.log('Deleting message')
          })
      },
      edit() {
        this.editing = true
        this.editedBody = this.displayedBody
      },
      cancelEdit() {
        this.editing = false
        this.editedBody = null
      },
      doneEdit() {
        this.updateMessage({id: this.id, body: this.editedBody})
          .then(data => {
            console.log('Updating message')
            this.cancelEdit()
          })
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
        return fetch(`/api/messages/${id}`, {
          method: 'DELETE'
        })
        .then(response => response.json())
      },
      updateMessage(message) {
        return fetch(`/api/messages/${message.id}`, {
          method: 'PUT',
          body: JSON.stringify(message),
        })
        .then(response => response.json())
      },
      clearMessage() {
        this.newMessage = new Message();
      }
	  // 1-3. メッセージを編集しよう
      // ...
    }
  });
})();
