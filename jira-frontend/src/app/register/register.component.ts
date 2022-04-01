import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  registerForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    firstname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    lastname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    email: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required, Validators.minLength(3)]),
  });

  constructor(private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
  }

  submit(){
    let body = {
      username: this.registerForm.get('username')?.value,
      firstname:this.registerForm.get('firstname')?.value,
      lastname: this.registerForm.get('lastname')?.value,
      password: this.registerForm.get('password')?.value,
      email: this.registerForm.get('email')?.value
    } 

    this.apiService.register(body)
    .subscribe(res => {
          console.log(res);
    })
    
  }

}
