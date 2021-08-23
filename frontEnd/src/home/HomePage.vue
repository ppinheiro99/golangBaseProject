<template>
    <div>
        <h1>Hi {{ user.email }}!</h1>
        <h3>Users: </h3>
        <ul v-if="users">
            <li v-for="user in users" :key="user.ID">
                {{user.ID + ' ' + user.email}}
                <span v-if="user.deleting"><em> - Deleting...</em></span>
                <span v-else-if="user.deleteError" class="text-danger"> - ERROR: {{user.deleteError}}</span>
                <span to="/HomePage" v-else> - <a @click="deleteUser(user.ID)" class="text-danger">Delete</a></span>
            </li>
        </ul>
        <p>
            <router-link to="/login">Logout</router-link>
        </p>
    </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import { userService } from '../_services/user.service'

export default {
     //console.warn(localStorage.getItem('user'))
      data () {
        return {
            user: JSON.parse(localStorage.getItem('user')),
            account: state => state.account,
            users: JSON.parse(localStorage.getItem('getAllusers'))
        }
    },
    created () {
        userService.getAll()
    },
    methods: {
            deleteUser(id){
                console.log(id)
                userService.delete(id)
            }
    }
};
</script>