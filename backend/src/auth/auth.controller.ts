import { Controller, Post, Body, HttpCode, HttpStatus } from '@nestjs/common';
import { AuthService } from './auth.service';

@Controller('api/auth')
export class AuthController {
  constructor(private authService: AuthService) {}

  @Post('login')
  @HttpCode(HttpStatus.OK)
  async login(@Body() body: { username: string; password: string }) {
    return this.authService.login(body.username, body.password);
  }

  @Post('register')
  async register(
    @Body() body: { username: string; password: string; name: string },
  ) {
    return this.authService.register(
      body.username,
      body.password,
      body.name,
    );
  }
}
