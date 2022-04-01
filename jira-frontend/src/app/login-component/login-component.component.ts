import { Component, OnInit } from '@angular/core';

import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-login-component',
  templateUrl: './login-component.component.html',
  styleUrls: ['./login-component.component.scss']
})
export class LoginComponentComponent implements OnInit {
  
  newLoginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
  }

  submit(){
    let body = {
      email:this.newLoginForm.get('email')?.value,
      password: this.newLoginForm.get('password')?.value,
    } 

    this.apiService.login(body)
    .subscribe(res => {
          console.log(res);
    })
    
  }

}
