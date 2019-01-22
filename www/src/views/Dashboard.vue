<template>
    <v-container grid-list-xl text-xs-center >
        <v-layout justify-center align-center column>
            <v-flex >
                <v-img :src="require('../assets/logo-plamienok.png')"
                        class="my-1"
                        contain
                        width="300"
                ></v-img>
            </v-flex>

            <v-flex v-if="!$store.state.logged" ma-3>
                <h2 class="display-1">
                    LOGIN
                </h2>
                <v-card min-width="400">
                    <v-alert v-model="showAlert" type="error" prominent :dismissible="true">{{alertMessage}}</v-alert>
                    <v-form v-model="valid" lazy-validation ref="form" class="ma-2">
                        <v-text-field label="Email" v-model="email" :rules="emailRules" required></v-text-field>
                        <v-text-field
                                label="Heslo"
                                v-model="password"
                                required
                                :append-icon="show1 ? 'visibility_off' : 'visibility'"
                                :type="show1 ? 'text' : 'password'"
                                name="input-10-1"
                                counter
                                @click:append="show1 = !show1">
                        </v-text-field>
                        <v-btn color="success" @click="login">Prihlásiť</v-btn>
                        <v-btn color="info" @click="forgot">Zabudnuté heslo</v-btn>
                    </v-form>
                </v-card>
            </v-flex>
        </v-layout>

        <v-layout justify-center>
            <v-flex class="ma-3 xs12 sm12 md6 lg6 ma-3">
                <h3 class="display-1">
                    Plamienok App
                </h3>
                <p class="subheading font-weight-regular">
                    Interná aplikácia pre neziskovú organizáciu Plamienok
                    <br>
                    <a href="https://www.plamienok.sk" target="_blank">www.plamienok.sk</a>
                </p>
            </v-flex>
        </v-layout>
        <v-layout justify-center>
            <v-flex class="ma-3 xs12 sm12 md6 lg6">
                <h3 class="display-1">Štatistika</h3>
                <v-layout justify-center>
                    <v-flex xs12 sm12 md12 lg12>
                        <v-card min-width="300">
                            <v-card-title><h4>Štatistika</h4></v-card-title>
                            <v-divider></v-divider>
                            <v-list dense>
                                <v-list-tile v-for="stat in stats" :key="stat.text">
                                    <v-list-tile-content>{{stat.text}}:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{stat.count}}</v-list-tile-content>
                                </v-list-tile>
                            </v-list>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-flex>
        </v-layout>

        <v-layout justify-center>
            <v-flex class="ma-3 xs12 sm12 md6 lg6">
                <h3 class="display-1">Aplikácia</h3>

                <p>
                    Aplikácia je Open Source aplikácia.
                    <v-btn to="/about">Viac o aplikácii</v-btn>
                </p>

            </v-flex>
        </v-layout>
    </v-container>
</template>


<script>
    export default {
        name: "Dashboard",

        data: () => ({
            valid: false,
            email: '',
            emailRules: [
                v => !!v || 'Email je vyžadovaný',
                v => /.+@.+/.test(v) || 'Email musí mať správny formát'
            ],
            password: '',
            stats: [
                {
                    text: 'Počet užívateľov',
                    count: 21
                },
                {
                    text: 'Počet klientov',
                    count: 12
                },
                {
                    text: 'Výjazdy za minulý týždeň',
                    count: 13
                },
            ],
            show1: false,
            alertMessage: '',
            showAlert: false,
        }),

        methods: {
            login: function() {
                this.showAlert=false;
                this.alertMessage='';

                if (this.valid && this.email !== "" && this.password !== "") {
                    this.$axios.post('/api/login',{
                            email:this.email,
                            password:this.password,
                        }).then(response => {
                            console.log('OK',response);
                            if (response.status == 202) {
                                this.showAlert=true;
                                this.alertMessage='Chyba: ' + response.data.Error;
                            } else {
                                this.$store.commit("login", {
                                    name: response.data.User.Name,
                                    email: response.data.User.Email,
                                });
                                localStorage.token = response.data.Token;
                                this.$axios.defaults.headers.common["token"] = localStorage.token;
                            }
                        }).catch(response => {
                            console.log('WRONG',response);
                            this.showAlert=true;
                            this.alertMessage='Chyba: ' + response;

                    });
                } else {
                    this.showAlert=true;
                    this.alertMessage='Chýbajú údaje na prihlásenie';
                }
            },

            forgot: function() {
                if (this.email !== "") {
                    console.log("Forgot", this.email);
                } else {
                    alert("Email musí byť zadaný");
                }
            },
        }
    }
</script>

<style scoped>

</style>