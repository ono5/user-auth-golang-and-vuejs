//  Home.vue
<template>
  <div class="container">
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
  name: "Home",
  setup() {
    const message = ref('You are not logged in!')

    onMounted(async () => {
      // user情報を取得
      // ログイン情報は、Cookieに保存してあるので、
      // リクエストするだけでOK
      const { data } = await axios.get('user')

      message.value = `Hi ${data.first_name} ${data.last_name}`
    })

    return {
      message
    }
  }

}
</script>

