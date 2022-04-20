import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import {Md5} from 'ts-md5/dist/md5'

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent implements OnInit {
  userProfileForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    firstname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    lastname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    email: new FormControl('', [Validators.required, Validators.minLength(3)])  
  });
  username = ""
  fullname = ""
  email = ""

  changePasswordForm = new FormGroup({
    oldPassword: new FormControl('', [Validators.required]),
    newPassword: new FormControl('', [Validators.required])
  })

  constructor(private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    this.apiService.getUserProfile()
    .subscribe((resp:any) => {
      this.userProfileForm.controls['username'].setValue(resp['resp']['username']);
      this.userProfileForm.controls['firstname'].setValue(resp['resp']['firstname']);
      this.userProfileForm.controls['lastname'].setValue(resp['resp']['lastname']);
      this.userProfileForm.controls['email'].setValue(resp['resp']['email_id']);
      this.username = resp['resp']['username']
      this.fullname = resp['resp']['firstname'] + " " + resp['resp']['lastname']
      this.email = resp['resp']['email_id']

    })
    

  }

  submit() {
    let body = {
      username: this.userProfileForm.get('username')?.value,
      firstname:this.userProfileForm.get('firstname')?.value,
      lastname: this.userProfileForm.get('lastname')?.value,
      email_id: this.userProfileForm.get('email')?.value
    } 

    this.apiService.editUserProfile(body)
    .subscribe(res => {
          console.log(res);
          this.username = res['resp']['username']
          this.fullname = res['resp']['firstname'] + " " + res['resp']['lastname']
          this.email = res['resp']['email_id']
    })
  }

  changePassword() {
    const md5 = new Md5()
    
    let old_password =  this.changePasswordForm.get('oldPassword')?.value
    let oldHashedPassword = md5.appendStr(old_password).end()

    let new_password =  this.changePasswordForm.get('newPassword')?.value
    let newHashedPassword = md5.appendStr(new_password).end()

    let body = {
      old_password: oldHashedPassword,
      new_password: this.changePasswordForm.get('newPassword')?.value
    } 

    this.apiService.changePassword(body)
    .subscribe(res => {
          console.log(res)
          
          alert('Password Cnged successfully!')
    })

  }

}
