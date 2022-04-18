import { Component, OnInit } from '@angular/core';

import { FormControl, FormGroup, Validators } from '@angular/forms';
import { has } from 'lodash';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import {Md5} from 'ts-md5/dist/md5'


@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent implements OnInit {

  newLoginForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(private apiService: ApiInterfaceService, private router: Router) { }

  ngOnInit(): void {
  }

  submit(){

    let password =  this.newLoginForm.get('password')?.value
    const md5 = new Md5()
    let hashedPassword = md5.appendStr(password).end()
    let body = {
      username:this.newLoginForm.get('username')?.value,
      password:hashedPassword
    } 

   
  }

}
