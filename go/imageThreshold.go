func imageThreshold(data[] byte, width int, height int)
{
array:= make([]int, widthheight)
s := 8
s2 := s / 2
t := 15
t2 := float64(100-t) / 100.0
for i := 0;
  i < width;
  i++
  {
  sum:= 0
for j := 0;
    j < height;
    j++
    {
    index:
      = jwidth + i
                     r : = int(data[index4 + 0])
                             g : = int(data[index4 + 1])
                                     b : = int(data[index4 + 2])
                                             data[index4] = byte(0.2126float64(r) + 0.7152float64(g) + 0.0722float64(b))
                                                 sum += int(data[index4]) if i == 0
      {
        array[index] = sum
      }
      else
      {
        array[index] = array[index - 1] + sum
      }
    }
  }
for
i:
  = 0;
i < width;
i++
{
for
j:
  = 0;
j < height;
j++
{
x1:
  = i - s2
            x2 : = i + s2
                           x1_1 : = x1 - 1 y1 : = j - s2
                                                          y2 : = j + s2
                                                                         y1_1 : = y1 - 1 if x1 < 0
  {
    x1 = 0
  }
  if x2
    >= width
    {
      x2 = width - 1
    }
  if x1_1
    < 0
    {
      x1_1 = 0
    }
  if y1
    < 0
    {
      y1 = 0
    }
  if y2
    >= height
    {
      y2 = height - 1
    }
  if y1_1
    < 0 {
          y1_1 = 0} count : = (x2 - x1) * (y2 - y1) index : = jwidth + i index1 : = y2width + x2 index2 : = y1_1width + x2 index3 : = y2width + x1_1 index4 : = y1_1width + x1_1 sum : = array[index1] - array[index2] - array[index3] + array[index4] if int(data[index4]) count <= int(sumt2)
    {
      data[index4 + 0] = 0 data[index4 + 1] = 0 data[index4 + 2] = 0
    }
  else
  {
    data[index4 + 0] = 255 data[index4 + 1] = 255 data[index4 + 2] = 255
  }
}
}
}