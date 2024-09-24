import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterOutlet } from '@angular/router';
import { lastValueFrom } from 'rxjs';

interface IUser {
  id: number
  username: string
  password: string
  createdAt: any
}

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    FormsModule
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit{
  public title = 'client'
  public username = ''
  public password = ''
  public users: IUser[] = []

  
  constructor(
    private httpClient: HttpClient
  ) {}

  async ngOnInit(){
    await this.loadNewsUsers()
  }

  async loadNewsUsers() {
    this.users = await lastValueFrom(this.httpClient.get<IUser[]>('/api/users'))
  }

  async addUser() {
    await lastValueFrom(this.httpClient.post('/api/users', {
      username: this.username,
      password: this.password
    }))

    await this.loadNewsUsers()
    
    this.username = ''
    this.password = ''
  }

}
