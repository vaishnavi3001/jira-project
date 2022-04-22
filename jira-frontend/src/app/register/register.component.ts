import { ThisReceiver } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { Md5 } from 'ts-md5/dist/md5'
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  inviteLink: string | null = null
  respMsg: string | null = null
  errorMsg: string | null = null
  registerForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    firstname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    lastname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    email: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required, Validators.minLength(3)]),
  });

  constructor(private apiService: ApiInterfaceService, private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
    this.inviteLink = this.route.snapshot.queryParams["invite-link"];
    console.log('register', this.inviteLink)
  }

  submit() {
    let password = this.registerForm.get('password')?.value
    const md5 = new Md5()
    let hashedPassword = md5.appendStr(password).end()
    let body = {
      username: this.registerForm.get('username')?.value,
      firstname: this.registerForm.get('firstname')?.value,
      lastname: this.registerForm.get('lastname')?.value,
      password: hashedPassword,
      email_id: this.registerForm.get('email')?.value
    }
    this.errorMsg = null;
    this.respMsg = null;
    this.apiService.register(body)
      .subscribe((res) => {
        if (res["status"]) {
          this.respMsg = "Registeration Successful"
        }
        else if(res["message"] == "USER_ALREADY_EXISTS") {
          this.errorMsg = "User with this username/password already exists in our system"
        }
      },(err)=>{
        this.errorMsg = "Error while registering. Please try again"
      })
  }

  navigateToLogin() {
    if (this.inviteLink) {
      this.router.navigateByUrl('login?invite-link=' + this.inviteLink)
    }
    else {
      this.router.navigate(['login'])
    }
  }

  isError(){
    return this.errorMsg && this.errorMsg.length!=0
  }

  isSuccess(){
    return this.respMsg && this.respMsg.length != 0
  }
}
