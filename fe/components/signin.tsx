"use client"

import { Icons } from "@/components/icons"
import { Button } from "./button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./card"
import { Input } from "./input"
import { Label } from "./label"

export function Signin() {
  return (
    <Card className="content w-30" >
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl">Sign In!</CardTitle>
        <CardDescription>
          Enter your email below to sign in!
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="grid gap-2">
          <Label htmlFor="email">Email</Label>
          <Input id="email" type="email" placeholder="m@example.com" />
        </div>
        <div className="grid gap-2">
          <Label htmlFor="password">Password</Label>
          <Input id="password" type="password" />
        </div>
      </CardContent>
      <CardFooter>
        <Button className="w-lg">Create account</Button>
      </CardFooter>
    </Card>
  )
}