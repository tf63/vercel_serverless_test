# https://zenn.dev/toppy/articles/04a55453706f9b
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World From Fast API"}