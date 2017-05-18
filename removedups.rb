#!/usr/bin/ruby

hf = ENV["HOME"] + "/.bash_history"
buf = File.open(hf).read

File.open(hf, "w") do |file|
    buf.lines.uniq.each {|line|
        file.write(line.strip + "\n")
    }
end


