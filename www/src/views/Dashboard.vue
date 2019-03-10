<template>
  <v-container
    grid-list-xl
    text-xs-center
  >
    <v-layout
      justify-center
      align-center
      column
    >
      <v-flex>
        <v-img
          :src="require('../assets/logo-plamienok.png')"
          class="my-1"
          contain
          width="300"
        />
      </v-flex>

      <v-flex
        v-if="!$store.state.logged"
        ma-3
      >
        <h2 class="display-1">
          LOGIN
        </h2>
        <v-card
          min-width="400"
          v-if="!forgotMode"
        >
          <v-alert
            v-model="showAlert"
            type="error"
            prominent
            :dismissible="true"
          >
            {{ alertMessage }}
          </v-alert>
          <v-form
            v-model="valid"
            lazy-validation
            ref="form"
            class="ma-2"
          >
            <v-text-field
              label="Email"
              v-model="email"
              :rules="emailRules"
              required
            />
            <v-text-field
              label="Heslo"
              v-model="password"
              required
              :append-icon="show1 ? 'visibility_off' : 'visibility'"
              :type="show1 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show1 = !show1"
            />
            <v-btn
              color="success"
              @click="login"
            >
              Prihlásiť
            </v-btn>
            <v-btn
              color="info"
              @click="forgot"
            >
              Zabudnuté heslo
            </v-btn>
          </v-form>
        </v-card>

        <v-card
          min-width="400"
          v-if="forgotMode"
        >
          <v-alert
            v-model="showAlert"
            type="error"
            prominent
            :dismissible="true"
          >
            {{ alertMessage }}
          </v-alert>
          <v-form
            v-model="valid"
            lazy-validation
            ref="formCode"
            class="ma-2"
          >
            <v-text-field
              label="Kód z emailu"
              v-model="code6"
              :rules="code6Rules"
              required
              mask="######"
            />

            <h2>Nové heslo (nepovinné)</h2>

            <v-text-field
              label="Heslo"
              v-model="password"
              :append-icon="show2 ? 'visibility_off' : 'visibility'"
              :type="show2 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show2 = !show2"
            />

            <v-text-field
              label="Heslo"
              v-model="password2"
              :append-icon="show3 ? 'visibility_off' : 'visibility'"
              :type="show3 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show3 = !show3"
            />

            <v-btn
              color="success"
              @click="sendCode6"
            >
              Poslať
            </v-btn>
            <v-btn
              color="info"
              @click="back"
            >
              Späť
            </v-btn>
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
          <a
            href="https://www.plamienok.sk"
            target="_blank"
          >
            www.plamienok.sk
          </a>
        </p>
      </v-flex>
    </v-layout>
    <v-layout justify-center>
      <v-flex class="ma-3 xs12 sm12 md6 lg6">
        <h3 class="display-1">
          Štatistika
        </h3>
        <v-layout justify-center>
          <v-flex
            xs12
            sm12
            md12
            lg12
          >
            <v-card min-width="300">
              <v-card-title><h4>Štatistika</h4></v-card-title>
              <v-divider />
              <v-list dense>
                <v-list-tile
                  v-for="stat in stats"
                  :key="stat.Text"
                >
                  <v-list-tile-content>{{ stat.Text }}:</v-list-tile-content>
                  <v-list-tile-content class="align-end">
                    {{ stat.Value }}
                  </v-list-tile-content>
                </v-list-tile>
              </v-list>
            </v-card>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>

    <v-layout justify-center>
      <v-flex class="ma-3 xs12 sm12 md6 lg6">
        <h3 class="display-1">
          Aplikácia
        </h3>

        <p>
          Aplikácia je Open Source aplikácia.
          <v-btn to="/about">
            Viac o aplikácii
          </v-btn>
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
            password2: '',
            stats: [],
            forgotMode: false,
            show1: false,
            show2: false,
            show3: false,
            alertMessage: '',
            showAlert: false,
            code6: '',
           code6Rules: [
                v => !!v || 'Kód musí byť zadaný',
                v => /^[0-9]{6}$/.test(v) || 'Kód je 6 číselných znakov'
            ],

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
                              this.loginSave(response.data);
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
                this.showAlert=false;
                this.alertMessage='';
                if (this.email !== "") {
                    console.log("Forgot", this.email);
                    this.$axios.post('/api/forgot',{email:this.email}).then(response => {
                        if (response.status == 202) {
                            this.showAlert=true;
                            this.alertMessage='Chyba: ' + response.data.Error;
                        } else {
                            this.password="";
                            this.password2="";
                            this.code6="";
                            this.forgotMode = true;
                        }
                    }).catch(response => {
                        console.log('WORNG',response)
                        this.showAlert=true;
                        this.alertMessage='Chyba: ' + response;
                    });
                } else {
                    this.alertMessage = "Email musí byť zadaný";
                    this.showAlert = true;
                }
            },

            sendCode6: function() {
                this.showAlert=false;
                this.alertMessage='';
                if (this.code6 === "") {
                  this.alertMessage = "Kód musí byť zadaný";
                  this.showAlert = true;
                } else if (this.password !== this.password2) {
                  this.alertMessage = "Opakované heslo nie je rovnaké";
                  this.showAlert = true;
                } else if (this.password.length > 0 && this.password.length < 6) {
                  this.alertMessage = "Heslo je príliš krátke";
                  this.showAlert = true;
                } else {
                    this.$axios.post('/api/checkCode6',{email:this.email, code6:this.code6, password: this.password}).then(response => {
                        if (response.status == 202) {
                            this.showAlert=true;
                            this.alertMessage='Chyba: ' + response.data.Error;
                        } else {
                            if (response.status == 202) {
                              this.showAlert=true;
                              this.alertMessage='Chyba: ' + response.data.Error;
                              this.forgotMode = true;
                            } else {
                              this.loginSave(response.data);
                            }
                        }
                    }).catch(response => {
                        console.log('WORNG',response)
                        this.showAlert=true;
                        this.alertMessage='Chyba: ' + response;
                    });
                }
            },

            loginSave: function(data) {
              console.log('loginSave', data.User)
              this.$store.commit("login", {
                name: data.User.Name,
                id: data.User.ID,
                email: data.User.Email
              });
              localStorage.token = data.Token;
              this.$axios.defaults.headers.common["token"] = localStorage.token;
              this.forgotMode = false;
            },

            back: function () {
                this.forgotMode = false;
            }
        },

        mounted: function() {
          console.log("mounted 2");
          this.$axios.get("/api/stats").then(response => {
                    console.log(response);
                    this.stats = response.data;
                  }
          ).catch(response => {
            console.log("Chyba",response);
          });
        }
    }
</script>

<style scoped>

</style>