import { Component, OnInit } from '@angular/core';

import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { Md5 } from 'ts-md5/dist/md5'

@Component({
  selector: 'app-login-component',
  templateUrl: './login-component.component.html',
  styleUrls: ['./login-component.component.scss']
})
export class LoginComponentComponent implements OnInit {
  inviteLink: string | null = null
  respMsg: string | null = null
  errorMsg: string | null = null
  newLoginForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(private apiService: ApiInterfaceService, private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.inviteLink = this.route.snapshot.queryParams["invite-link"];
    console.log("inviteLink", this.inviteLink)
  }

  submit() {
    let password = this.newLoginForm.get('password')?.value
    const md5 = new Md5()
    let hashedPassword = md5.appendStr(password).end()
    let body = {
      username: this.newLoginForm.get('username')?.value,
      password: hashedPassword
    }
     
    this.errorMsg = null;
    this.apiService.login(body)
      .subscribe(res => {
        // console.log(res["resp"]["access_token"])
        if (res["status"]) {
          this.apiService.setToken(res["resp"]["access_token"]);
          if (this.inviteLink != null) {
            this.router.navigateByUrl('/home/join-project/' + this.inviteLink)
          } else {
            this.router.navigateByUrl("/home/projects")
          }
        } else {
          this.errorMsg = "Invalid Username/Password"
        }
      },(err)=>{
        console.log("error");
      })
  }

  navigateToRegister() {
    if (this.inviteLink) {
      this.router.navigateByUrl("/register?invite-link=" + this.inviteLink)
    }
    else {
      this.router.navigate(["register"])
    }
  }

  isError(){
    return this.errorMsg && this.errorMsg.length!=0
  }

  isSuccess(){
    return this.respMsg && this.respMsg.length != 0
  }
}
